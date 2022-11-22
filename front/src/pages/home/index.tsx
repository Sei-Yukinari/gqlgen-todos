import {FC, JSXElementConstructor, Key, ReactElement, ReactFragment, ReactPortal, useEffect} from "react";
import {useTodo} from "@/hooks/useTodo";
import {useQuery, useSubscription} from "urql";
import {gql} from "graphql-request";

const HomePage: FC = () => {

    const todo = useTodo()

    useEffect(() => {
        (async () => {
            // await todo.createTodo()
            await todo.findTodos()
            // await todo.messagePosted()
        })()
    }, []);

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
    const newMessages = gql`
        subscription($user: String!) {
          messagePosted(user: $user) {
            id
            user
            text
            createdAt
          }
        }
`;

    const handleSubscription = (messages = [], response: {
        messagePosted: any;
        data: any; }) => {
        console.log(99999999,response.messagePosted.text)
        return [response.messagePosted, ...messages];
    };
    const [res] = useSubscription({query: newMessages, variables: {user: 'usercccc'}}, handleSubscription);
    if (!res.data) {
        return <p>No new messages</p>;
    }

    console.log(7777777,res)
    const {data, fetching, error} = result;
    console.log(data)
    if (fetching) return <p>Loading...</p>;
    if (error) return <p>Oh no... {error.message}</p>;




    return (
        <>
            <h1>Hello World!</h1>
            <ul>
                {data.todos.map((todo: { id: Key | null | undefined; text: string | number | boolean | ReactElement<any, string | JSXElementConstructor<any>> | ReactFragment | ReactPortal | null | undefined; }) => (
                    <li key={todo.id}>{todo.text}</li>
                ))}
            </ul>
        </>
    )
}

export default HomePage