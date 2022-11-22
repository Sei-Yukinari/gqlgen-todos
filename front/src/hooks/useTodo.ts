import {GraphQLClient, gql} from 'graphql-request'
import {createClient, useMutation, useQuery} from 'urql';

export const useTodo = () => {
    const [result, reexecuteQuery] = useQuery({
        query: gql`
                query findTodos {
                  todos {
                    text
                    done
                    user {
                      name
                    }
                  }
                }
  `,
    });
    const [updateTodoResult, updateTodo] = useMutation(gql`
            mutation createTodo {
              createTodo(input: { text: "todo", userId: "1" }) {
                user {
                  id
                }
                text
                done
              }
            }
  `)
    const findTodos = async () => {

    }

    const createTodo = async () => {
        updateTodo().then(result => {
            if (result.error) {
                console.error('Oh no!', result.error);
            }
            console.log(JSON.stringify(result.data, undefined, 2))
        })
    }


    const messagePosted = () => {
    }

    return {findTodos, createTodo, messagePosted}
}