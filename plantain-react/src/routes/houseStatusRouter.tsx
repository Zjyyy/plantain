import HouseStatusPanel  from "src/sections/houseStatus/HouseStatusPanel";
import ApartmentStatusPanel from "src/sections/houseStatus/ApartmentStatusPanel";

const houseStatusRouter = [
    {
        path:'/housePanel',
        element:<HouseStatusPanel/>
    },
    {
        path:'/house/:id',
        element:<ApartmentStatusPanel/>
    }
];

export default houseStatusRouter;