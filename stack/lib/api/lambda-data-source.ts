import { Construct, Duration } from '@aws-cdk/core'
import * as lambda from '@aws-cdk/aws-lambda'
import * as iam from '@aws-cdk/aws-iam'
import {Props} from "../props";
import {Name} from '../../utils/resource-name'

export const LambdaDataSource = (scope: Construct, props: Props): lambda.IFunction => {
    const functionName = Name(props, 'graphql-datasource')

    const environment = {
        GEOFF_CDK_NOTES_TABLE_NAME: props.staticConfig.geoffCdkNotesTableName
    }

    const fn = new lambda.Function(scope, functionName, {
        functionName,
        description: 'CDK NOTES API datasource',
        runtime: lambda.Runtime.GO_1_X,
        timeout: Duration.seconds(30),
        code: lambda.Code.fromAsset('../bin/lambda', {exclude: ['**', '!appsync']}),
        memorySize: 1024,
        handler: 'appsync',
        tracing: lambda.Tracing.ACTIVE,
        environment: environment,
    })

    fn.addToRolePolicy(
        new iam.PolicyStatement({
            effect: iam.Effect.ALLOW,
            actions: [
                'dynamodb:DescribeTable',
                'dynamodb:Query',
                'dynamodb:Scan',
                'dynamodb:GetItem',
                'dynamodb:PutItem',
                'dynamodb:UpdateItem',
                'dynamodb:DeleteItem',
                'dynamodb:BatchWriteItem',
            ],
            resources: [
                `arn:aws:dynamodb:*:*:table/${props.staticConfig.geoffCdkNotesTableName}`,
                `arn:aws:dynamodb:*:*:table/${props.staticConfig.geoffCdkNotesTableName}/index/*`
            ],
        }),
    )

    return fn
}