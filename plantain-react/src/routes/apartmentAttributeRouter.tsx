import { ApartmentAttribute } from "src/sections/apartmentAttribute";
import MultimediaMaterialsAttribute from "src/sections/apartmentAttribute/MultimediaMaterialsAttribute";

const apartmentAttributeRouter = [
    {
        path:'/apartmentAttribute',
        element:<ApartmentAttribute/>
    },
    {
        path:'/apartmentAttribute/multimediaMaterialsAttribute',
        element:<MultimediaMaterialsAttribute/>
    }
]

export default apartmentAttributeRouter;