import * as appSync from '@aws-cdk/aws-appsync'
import * as lambda from '@aws-cdk/aws-lambda'
import {CfnOutput, Construct, Duration, Expiration} from '@aws-cdk/core'
import {Name} from '../../utils/resource-name'
import {Props} from '../props'

export const AppSync = (stack: Construct, props: Props, lambdaDataSource: lambda.IFunction) => {
    const name = Name(props, 'GG')

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
            }
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

    const dataSource = api.addLambdaDataSource('lambdaDatasource', lambdaDataSource, {
        description: 'The lambda supporting the GraphQL API',
    })

    dataSource.createResolver({
        typeName: 'Mutation',
        fieldName: 'createNote',
        requestMappingTemplate: appSync.MappingTemplate.fromFile('templates/createNote.vtl'),
        responseMappingTemplate: appSync.MappingTemplate.fromFile('templates/to-json.vtl'),
    })

/*    dataSource.createResolver({
        typeName: 'Query',
        fieldName: 'listNotes',
        requestMappingTemplate: appSync.MappingTemplate.fromFile('templates/listNotes.vtl'),
        responseMappingTemplate: appSync.MappingTemplate.fromFile('templates/to-json.vtl'),
    })*/

    return api
}
