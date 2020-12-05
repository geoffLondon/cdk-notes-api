import {Construct} from '@aws-cdk/core'
import * as dynamodb from '@aws-cdk/aws-dynamodb'
import {Props} from "../props"

export const DynamodbServicesTable = (stack: Construct, props: Props): dynamodb.ITable => {
    return new dynamodb.Table(stack, props.staticConfig.geoffCdkNotesTableName, {
        tableName: props.staticConfig.geoffCdkNotesTableName,
        partitionKey: {name: 'pk', type: dynamodb.AttributeType.STRING},
        billingMode: dynamodb.BillingMode.PAY_PER_REQUEST,
    })
}
