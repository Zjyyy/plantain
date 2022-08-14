import { ClientList } from "src/sections/clientList";
import ClientMessage from "src/sections/clientList/ClientMessage";

const ClientListRouter = [
    {
        path:'/clientList',
        element:<ClientList/>
    },
    // {
    //     path:'/clientList/clientMessage',
    //     element:<ClientMessage/>
    // },
    {
        path:'/clientList/clientMessage/:id',
        element:<ClientMessage/>
    }
]

export default ClientListRouter;