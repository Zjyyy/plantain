import { makeAutoObservable } from 'mobx';
import ClientResourceAPI from 'src/apis/ClientResourceAPI';
import { AxiosResponse } from 'axios';
import errorMessageStore from 'src/stores/components/errorModalStore';
import { IAPIResult } from 'src/model/commonModel';
import { IClientMessage } from 'src/model/clientMessageModalModel'; 

export default class ClientMessageModalStore{
    clientResourceAPI:ClientResourceAPI = new ClientResourceAPI()

    clientMessage:IClientMessage = {
        name:"",
        profile:"",
        idNumber:"",
        phoneNumber:"",
        weChatNumber:"",
        rate:0,
        des:""
    }
    constructor(){
        makeAutoObservable(this)
    }
    initializeClientMessage(){
        this.clientMessage = {
            name:"",
            profile:"",
            idNumber:"",
            phoneNumber:"",
            weChatNumber:"",
            rate:0,
            des:""
        }
    }

    async submitClientMessage(){
        return await this.clientResourceAPI.postClientMessage(this.clientMessage).then((response:AxiosResponse)=>{
            const retrieved = response.data as IAPIResult
            if(retrieved.success){
                return true
            }
            else{
                errorMessageStore.Error("添加客户失败",retrieved.resultMessage)
                return false
            }
        })
        return false
    }

    onChangeName(val:string){
        this.clientMessage.name = val
    }
    onChangeIDNumber(val:string){
        this.clientMessage.idNumber = val
    }
    onChangeProfile(val:string){
        this.clientMessage.profile = val
    }
    onChangePhoneNumber(val:string){
        this.clientMessage.phoneNumber = val
    }
    onChangeWeChatNumber(val:string){
        this.clientMessage.weChatNumber = val
    }
    onChangeRate(val:number){
        this.clientMessage.rate = val
    }
    onChangeDes(val:string){
        this.clientMessage.des = val
    }
}