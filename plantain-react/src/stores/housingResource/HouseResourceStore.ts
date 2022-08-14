import { makeAutoObservable,runInAction } from 'mobx';
import HousingResourceAPI from 'src/apis/HousingResourceAPI';
import { AxiosResponse } from 'axios';
import errorMessageStore from 'src/stores/components/errorModalStore';
import { IAPIResult } from 'src/model/commonModel';
import { IHousePanel,IHouse, IApartmentPanel } from 'src/model/houseResourceModel';

export default class HouseResourceStore {
    houseApi:HousingResourceAPI = new HousingResourceAPI()
    housePanelList:IHousePanel[] = []
    apartmentPanelList:IApartmentPanel[] = []
    layout:any = []
    apartmentLayout:any = []
    house:IHouse = {
        id:'',
        name:'',
        des:'',
        rentedNumber:0,
        sumNumber:0,
        monthlyIncome:0
    }

    constructor(){
        makeAutoObservable(this)
    }

    initializeHouse(){
        this.house = {
            id:'',
            name:'',
            des:'',
            rentedNumber:0,
            sumNumber:0,
            monthlyIncome:0
        }
    }

    loadHouse(id:string){
        this.houseApi.loadHouseById(id).then((response:AxiosResponse)=>{
            const retrieved = response.data as IAPIResult
            if(retrieved.success){
                const retrievedHouse = retrieved.resultObject as IHouse
                if(retrievedHouse != null){
                    const item = retrievedHouse
                    this.house = {
                        id:item.id,
                        name:item.name,
                        des:item.des,
                        rentedNumber:item.rentedNumber,
                        sumNumber:item.sumNumber,
                        monthlyIncome:item.monthlyIncome
                    }
                }
                else{
                    this.initializeHouse()
                }
            }
            else{
                errorMessageStore.Error("加载房子信息错误",retrieved.resultMessage)
            }
        })
    }
    async removeHouse(id:string){
        return await this.houseApi.deleteHouseById(id).then((response:AxiosResponse)=>{
            const retrieved = response.data as IAPIResult
            if(retrieved.success){
                return true
            }
            else{
                errorMessageStore.Error("删除楼栋失败",retrieved.resultMessage)
                return false
            }
        })
        return false
    }
    loadApartmentList(){
        let x = 0
        let y = 0
        this.apartmentLayout = []
        this.houseApi.loadApartmentList().then((response:AxiosResponse)=>{
            const retrieved = response.data as IAPIResult
            if(retrieved.success){
                const retrievedApartmentList = retrieved.resultObject as IApartmentPanel[]
                if(retrievedApartmentList.length > 0){
                    const apartmentPanelList:IApartmentPanel[] = []
                    retrievedApartmentList.forEach((item:IApartmentPanel)=>{
                        const apartmentPanel:IApartmentPanel = {
                            id:item.id,
                            roomBelongsHouseId:item.roomBelongsHouseId,
                            apartmentName:item.apartmentName,
                            monthlyRent:item.monthlyRent,
                            modeOfPayment:item.modeOfPayment,
                            guaranteeDeposit:item.guaranteeDeposit,
                            collectRentDate:item.collectRentDate,
                            leaseTime:item.leaseTime
                        }
                        apartmentPanelList.push(apartmentPanel)
                        this.apartmentLayout.push({i:apartmentPanel.id,x:x,y:y,w:2,h:1,isResizable:false})
                        x += 3
                        if(x > 6){
                            x = 0
                            y += 1
                        }
                    })
                    this.apartmentPanelList = apartmentPanelList
                }
            }
            else{
                this.apartmentPanelList = []
                errorMessageStore.Error("加载房源信息错误",retrieved.resultMessage)
            }
        })
    }
    loadHouseList(){
        let x = 0
        let y = 0
        this.layout = []
        this.houseApi.loadHouseList().then((response:AxiosResponse)=>{
            const retrieved = response.data as IAPIResult
            if(retrieved.success){
                const retrievedHouseList = retrieved.resultObject as IHousePanel[]
                if(retrievedHouseList.length > 0){
                    const housePanelList:IHousePanel[] = []
                    retrievedHouseList.forEach((item:IHousePanel)=>{
                        const housePanel:IHousePanel = {
                            id:item.id,
                            name:item.name,
                            des:item.des,
                            rentedNumber:item.rentedNumber,
                            sumNumber:item.sumNumber,
                            monthlyIncome:item.monthlyIncome
                        }
                        housePanelList.push(housePanel)
                        this.layout.push({i:housePanel.id,x:x,y:y,w:2,h:1,isResizable:false})
                        x += 2
                        if(x > 6){
                            x = 0
                            y += 1
                        }
                    })
                    this.housePanelList = housePanelList
                    this.layout.map((item:any)=>{
                        console.warn(item.i + " " +item.x+" "+item.y)
                    })
                }
                else{
                    this.housePanelList = []
                }
            }
            else{
                this.housePanelList = []
                errorMessageStore.Error("加载房源信息错误",retrieved.resultMessage)
            }
        })
    }
}