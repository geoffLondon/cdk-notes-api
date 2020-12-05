import * as appSync from '@aws-cdk/aws-appsync'
import * as lambda from '@aws-cdk/aws-lambda'
import {CfnOutput, Construct, Duration, Expiration} from '@aws-cdk/core'
import * as cognito from '@aws-cdk/aws-cognito'
import {Name} from '../../utils/resource-name'
import {Props} from '../props'

export const AppSync = (stack: Construct, props: Props, lambdaDataSource: lambda.IFunction) => {
    const name = Name(props, 'api')

    const api = new appSync.GraphqlApi(stack, name, {
        name: name,
        logConfig: {
            fieldLogLevel: appSync.FieldLogLevel.ALL,
        },
        schema: appSync.Schema.fromAsset('../schema.graphql'),
        authorizationConfig: {
            defaultAuthorization: {
                authorizationType: appSync.AuthorizationType.API_KEY,
                apiKeyConfig: {
                    expires: Expiration.after(Duration.days(365))
                }
            },
            additionalAuthorizationModes: [{
                authorizationType: appSync.AuthorizationType.USER_POOL,
                userPoolConfig: {
                    userPool: cognito.UserPool.fromUserPoolId(stack, "smartnumber-user-pool", props.config.smartnumbersUserPoolId)
                }
            }]
        },
        xrayEnabled: true
    })

    // print out the AppSync GraphQL endpoint to the terminal
    new CfnOutput(stack, "GraphQLApiUrl", {
        value: api.graphqlUrl
    })

    // print out the AppSync API Key to the terminal
    new CfnOutput(stack, "GraphQLApiKey", {
        value: api.apiKey || ''
    })

/*    const dataSource = api.addLambdaDataSource('LAMBDA', lambdaDataSource, {
        description: 'The lambda supporting the GraphQL API',
    })

    dataSource.createResolver({
        typeName: 'Query',
        fieldName: 'getService',
        requestMappingTemplate: appSync.MappingTemplate.fromFile('templates/getService.vtl'),
        responseMappingTemplate: appSync.MappingTemplate.fromFile('templates/to-json.vtl'),
    })
    */
    return api
}
