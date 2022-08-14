import React from "react";
import { rootStore } from "./rootStore";

const storeContext = React.createContext(rootStore);

export default storeContext;