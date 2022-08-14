import commonAxiosRequests from '../utils/agent';

export default class ClientResourceAPI{
    public loadClientMessageList(){
        return commonAxiosRequests.get(
            '/api/ClientResource/Client/List'
        );
    }
    public loadClientMessage(id:string){
        return commonAxiosRequests.get(
            '/api/ClientResource/Client' +
            '/' + id
        );
    }
    public postClientMessage(requestBody:{}){
        return commonAxiosRequests.post(
            '/api/ClientResource/Client',
            requestBody
        );
    }
    public putClientMessage(requestBody:{}){
        return commonAxiosRequests.put(
            '/api/ClientResource/Client',
            requestBody
        );
    }
    public deleteClientMessage(id:string){
        return commonAxiosRequests.del(
            '/api/ClientResource/Client' +
            '/' + id
        );
    }
}