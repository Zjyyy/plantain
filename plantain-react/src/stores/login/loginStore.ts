import { makeAutoObservable } from "mobx";

export default class LoginStore{
    constructor(){
        makeAutoObservable(this);
    }
}