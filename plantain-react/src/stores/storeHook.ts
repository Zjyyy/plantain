import React,{ useContext } from "react";
import StoreContext from './storeContext';
import { observer } from 'mobx-react-lite';

const useStores = () => React.useContext(StoreContext);

export {
    observer,
    useStores
}