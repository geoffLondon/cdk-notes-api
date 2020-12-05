import { Construct } from '@aws-cdk/core'
import { Props } from '../props'
import * as iam from '@aws-cdk/aws-iam'
import { Effect } from '@aws-cdk/aws-iam'
import { Name } from '../../utils/resource-name'

export const IamRoleAppSync = (stack: Construct, props: Props, lambdaArn: string): iam.Role => {
    const roleId = Name(props, 'appsync-role')

    const role = new iam.Role(stack, roleId, {
        assumedBy: new iam.ServicePrincipal('appsync.amazonaws.com'),
        description: 'App Sync service role',
    })

    role.addToPolicy(
        new iam.PolicyStatement({
            effect: Effect.ALLOW,
            actions: ['lambda:invokeFunction', 'logs:CreateLogGroup', 'logs:CreateLogStream', 'logs:PutLogEvents'],
            resources: [lambdaArn],
        })
    )

    return role
}
