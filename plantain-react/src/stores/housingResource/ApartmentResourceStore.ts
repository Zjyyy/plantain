import { makeAutoObservable,runInAction } from 'mobx';
import HousingResourceAPI from 'src/apis/HousingResourceAPI';
import { AxiosResponse } from 'axios';
import errorMessageStore from 'src/stores/components/errorModalStore';
import { IAPIResult } from 'src/model/commonModel';
import { IHousePanel,IHouse } from 'src/model/houseResourceModel';

export default class ApartmentResourceStore{
    apartmentResourceAPI:HousingResourceAPI = new HousingResourceAPI()
    
    constructor(){
        makeAutoObservable(this)
    }
}