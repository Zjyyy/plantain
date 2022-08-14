import { RouteObject } from "react-router-dom";
import { Workbench } from "../sections/workbench";
import houseStatusRouter from "./houseStatusRouter";
import apartmentAttributeRouter from "./apartmentAttributeRouter"
import clientListRouter from "./clientListRouter";
const routes:RouteObject[] = [
    {
        path:'/workbench',
        element:<Workbench/>
    },
    {
        path:'/',
        element:<Workbench/>
    },
    ...houseStatusRouter,
    ...apartmentAttributeRouter,
    ...clientListRouter
];

export default routes;