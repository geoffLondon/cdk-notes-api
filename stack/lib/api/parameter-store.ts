import {Construct} from '@aws-cdk/core'
import {Props} from '../props'
import * as ssm from '@aws-cdk/aws-ssm'

export const ParameterStore = (stack: Construct, props: Props, graphqlUrl: string) => {

    const serviceParams = JSON.stringify({
        version: '2020.11.16',
        graphQlEndpoint: graphqlUrl
    })

    return new ssm.StringParameter(stack, 'Parameter', {
        parameterName: props.staticConfig.parameterName,
        description: 'CDK Notes Parameters',
        stringValue: serviceParams
    })

}