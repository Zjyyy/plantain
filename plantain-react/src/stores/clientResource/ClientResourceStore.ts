import { makeAutoObservable,runInAction } from 'mobx';
import ClientResourceAPI from 'src/apis/ClientResourceAPI';
import { AxiosResponse } from 'axios';
import errorMessageStore from 'src/stores/components/errorModalStore';
import { IAPIResult } from 'src/model/commonModel';
import { IClientMessage,IClientPanel } from 'src/model/clientResourceModel'; 

export default class ClientResourceStore{
    clientApi:ClientResourceAPI = new ClientResourceAPI()
    clientPanelList:IClientPanel[] = []
    clientMessage:IClientMessage = {
        id:"",
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
    async submitClientMessage(){
        return await this.clientApi.putClientMessage(this.clientMessage).then((response:AxiosResponse)=>{
            const retrieved = response.data as IAPIResult
            if(retrieved.success){
                return true
            }
            else{
                errorMessageStore.Error("保存客户失败",retrieved.resultMessage)
                return false
            }
        })
        return false
    }
    
    async removeClientMessage(id:string){
        return await this.clientApi.deleteClientMessage(id).then((response:AxiosResponse)=>{
            const retrieved = response.data as IAPIResult
            if(retrieved.success){
                return true
            }
            else{
                errorMessageStore.Error("删除客户信息失败",retrieved.resultMessage)
                return false
            }
        })
        return false
    }
    
    loadClientMessage(id:string){
        this.clientApi.loadClientMessage(id).then((response:AxiosResponse)=>{
            const retrieved = response.data as IAPIResult
            if(retrieved.success){
                const retrievedClientMessage = retrieved.resultObject as IClientMessage
                if(retrievedClientMessage != null){
                    const item = retrievedClientMessage
                    this.clientMessage = {
                        id:item.id,
                        name:item.name,
                        profile:item.profile,
                        idNumber:item.idNumber,
                        phoneNumber:item.phoneNumber,
                        weChatNumber:item.weChatNumber,
                        rate:item.rate,
                        des:item.des
                    }
                }
                else{
                    this.clientMessage = {
                        id:"",
                        name:"",
                        profile:"",
                        idNumber:"",
                        phoneNumber:"",
                        weChatNumber:"",
                        rate:0,
                        des:""
                    }
                }
            }
            else{
                errorMessageStore.Error("加载客户信息错误 ",retrieved.resultMessage)
            }
        })
    }

    loadClientPanelList(){
        this.clientApi.loadClientMessageList().then((response:AxiosResponse)=>{
            const retrieved = response.data as IAPIResult
            if(retrieved.success){
                const retrievedClientMessageList = retrieved.resultObject as IClientMessage[]
                if(retrievedClientMessageList.length > 0){
                    const clientPanelList:IClientPanel[] = [];
                    retrievedClientMessageList.forEach((item:IClientMessage)=>{
                        const clientPanel:IClientPanel = {
                            id:item.id,
                            name:item.name,
                            des:item.des,
                            profile:item.profile
                        }
                        clientPanelList.push(clientPanel)
                    })
                    this.clientPanelList = clientPanelList;
                    console.warn("test")
                }
                else{
                    this.clientPanelList = []
                }
            }
            else{
                this.clientPanelList = []
                errorMessageStore.Error("加载客户信息面板故障",retrieved.resultMessage)
            }
        })
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