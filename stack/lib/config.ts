import { DevelopmentStackConfig } from '../conf/development'

export interface StaticConfig {
    parameterName: string
    geoffCdkNotesTableName: string
}

export interface Config {
    stage: string
    environment: string
}

export const NewConfig = (stage: string): Config => {
    switch (stage) {
        case 'development':
            return DevelopmentStackConfig
        default:
            return DevelopmentStackConfig
    }
}
