import * as cdk from '@aws-cdk/core'
import { Props } from './props'
import { LambdaDataSource } from './api/lambda-data-source'
import { IamRoleAppSync } from './api/iam-role-app-sync'
import { IamRoleStack } from './api/iam-role-stack'
import { AppSync } from './api/app-sync'
import { DynamodbServicesTable } from './api/dynamodb-services-table'

export class Stack extends cdk.Stack {
    constructor(scope: cdk.Construct, id: string, props: Props) {
        super(scope, id, props)

        const lambdaDataSource = LambdaDataSource(this, props)

        IamRoleStack(this, props, lambdaDataSource.functionArn)
        IamRoleAppSync(this, props, lambdaDataSource.functionArn)

        AppSync(this, props, lambdaDataSource)

        DynamodbServicesTable(this, props)
    }
}
