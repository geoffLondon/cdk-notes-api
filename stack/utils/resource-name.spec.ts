import { Name } from './resource-name'
import { Props } from '../lib/props'
import { DevelopmentConfig } from '../conf/development'
import { StaticConfig } from '../conf/static'

describe('Name', () => {
    it('returns a name base on the stack and stage', () => {
        const props: Props = {
            config: DevelopmentConfig,
            staticConfig: StaticConfig,
            stackName: 'hello',
        }

        const name = Name(props, 'Molly')

        expect(name).toEqual('hello-Molly')
    })
})
