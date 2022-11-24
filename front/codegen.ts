import type { CodegenConfig } from '@graphql-codegen/cli'

const config: CodegenConfig = {
  overwrite: true,

  schema: {
    'http://localhost:8080/query': {
      headers: { Origin: 'http://localhost:3000' },
    },
  },
  documents: 'src/**/*.graphql',
  generates: {
    'src/gql/graphql.ts': {
      plugins: [
        'typescript',
        'typescript-operations',
        'typescript-urql',
      ],
    },
    './graphql.schema.json': {
      plugins: ['introspection'],
    },
  },
  hooks: {
    afterOneFileWrite: 'npx eslint --fix',
  },
}

export default config
