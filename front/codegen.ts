
import type { CodegenConfig } from '@graphql-codegen/cli';

const config: CodegenConfig = {
  overwrite: true,

  schema: "../api/graph/schema/*.graphql",
  documents: "src/**/*.graphql",
  generates: {
    "src/gql/graphql.ts": {
      plugins: [
          "typescript",
          "typescript-operations",
          "typescript-urql",
      ]
    },
    "./graphql.schema.json": {
      plugins: ["introspection"]
    }
  }
};

export default config;
