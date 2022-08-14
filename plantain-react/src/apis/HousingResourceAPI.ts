import commonAxiosRequests from '../utils/agent';

export default class HousingResourceAPI{
    public loadHouseList(){
        return commonAxiosRequests.get(
            '/api/HousingResource/Houses/List'
        );
    }
    public loadHouseById(id:string){
        return commonAxiosRequests.get(
            '/api/HousingResource/Houses' +
            '/' + id
        );
    }
    public postHouse(requestBody:{}){
        return commonAxiosRequests.post(
            '/api/HousingResource/Houses',
            requestBody
        );
    }
    public deleteHouseById(id:string){
        return commonAxiosRequests.del(
            '/api/HousingResource/Houses' +
            '/' + id
        );
    }
    public putHouse(requestBody:{}){
        return commonAxiosRequests.put(
            '/api/HousingResource/Houses',
            requestBody
        );
    }

    public patchHouse(id:string,requestBody:{}){
        return commonAxiosRequests.patch(
            '/api/HousingResource/Houses/'+id,
            requestBody
        );
    }

    public loadApartmentList(){
        return commonAxiosRequests.get(
            '/api/HousingResource/Apartment/List'
        );
    }
    public loadApartmentById(id:string){
        return commonAxiosRequests.get(
            '/api/HousingResource/Apartment' +
            '/' + id
        );
    }
    public postApartment(requestBody:{}){
        return commonAxiosRequests.post(
            '/api/HousingResource/Apartment',
            requestBody
        );
    }
    public deleteApartmentById(id:string){
        return commonAxiosRequests.del(
            '/api/HousingResource/Apartment' +
            '/' + id
        );
    }
    public putApartment(requestBody:{}){
        return commonAxiosRequests.post(
            '/api/HousingResource/Apartment',
            requestBody
        );
    }
}