import * as cdk from '@aws-cdk/core'
import { Props } from './props'
import { LambdaDataSource } from './api/lambda-data-source'
import { IamRoleAppSync } from './api/iam-role-app-sync'
import { IamRoleStack } from './api/iam-role-stack'
import { AppSync } from './api/app-sync'
import { ParameterStore } from './api/parameter-store'
import { DynamodbServicesTable } from './api/dynamodb-services-table'

export class Stack extends cdk.Stack {
  constructor(scope: cdk.Construct, id: string, props: Props) {
    super(scope, id, props)

    const lambdaDataSource = LambdaDataSource(this, props)

    IamRoleStack(this, props, lambdaDataSource.functionArn)
    IamRoleAppSync(this, props, lambdaDataSource.functionArn)

    const appSync = AppSync(this, props, lambdaDataSource)

    ParameterStore(this, props, appSync.graphqlUrl)

    DynamodbServicesTable(this, props)
  }
}
