import { makeAutoObservable } from 'mobx';
import HousingResourceAPI from 'src/apis/HousingResourceAPI';
import { AxiosResponse } from 'axios';
import errorMessageStore from 'src/stores/components/errorModalStore';
import { IAPIResult } from 'src/model/commonModel';
import { IHouse } from 'src/model/houseModalModel'; 

export default class HouseModalComponent {
    housingResourceAPI:HousingResourceAPI = new HousingResourceAPI()

    house:IHouse = {
        id:'',
        name:'',
        des:''
    }

    constructor(){
        makeAutoObservable(this)
    }

    initializeHouse(){
        this.house = {
            id:'',
            name:'',
            des:''
        }
    }
    async patchHouseMessage(id:string){
        return await this.housingResourceAPI.patchHouse(id,[
            {
                path:"/name",
                op:"replace",
                value:this.house.name
            },
            {
                path:"/des",
                op:"replace",
                value:this.house.des
            }
        ]).then()
    }
    async submitHouse(){
        return await this.housingResourceAPI.postHouse(this.house).then((response:AxiosResponse)=>{
            const retrieved = response.data as IAPIResult
            if(retrieved.success){
                return true
            }
            else{
                errorMessageStore.Error("添加楼栋失败",retrieved.resultMessage)
                return false
            }
        })
        return false
    }

    loadHouseMessage(houseId:string){
        return this.housingResourceAPI.loadHouseById(houseId).then((response:AxiosResponse)=>{
            const retrieved = response.data as IAPIResult
            if(retrieved.success){
                const retrievedHouseMessage = retrieved.resultObject
                if(retrievedHouseMessage != null){
                    const item = retrievedHouseMessage
                    this.house = {
                        id:item.id,
                        name:item.name,
                        des:item.des
                    }
                }
                else{
                    this.initializeHouse()
                }
            }
            
        })
    }

    onChangeName(val:string){
        this.house.name = val
    }
    onChangeDes(val:string){
        this.house.des = val
    }
}