#!/usr/bin/env node
import 'source-map-support/register'
import * as cdk from '@aws-cdk/core'
import { Stack } from '../lib/stack'
import { Props } from '../lib/props'
import { NewConfig } from '../lib/config'
import { StackStaticConfig } from '../conf/static'

const app = new cdk.App()

const stage = process.env.STAGE || 'development'
const team = 'personal'
const stackName = 'cdk-notes-api'
const description = 'A note taking application using AWS-CDK. (GG)'

const config = NewConfig(stage)
const staticConfig = StackStaticConfig

const props: Props = {
    tags: {
        environment: config.environment,
        stage: config.stage,
        stack: stackName,
        team,
    },
    description,
    stackName,
    config,
    staticConfig,
}

new Stack(app, stackName, props)
