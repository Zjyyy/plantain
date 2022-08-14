import  axios,{ AxiosError, AxiosResponse } from 'axios';
// import appConfig from '../configs/App.config';
import errorMessageStore from '../stores/components/errorModalStore';
import { getAuthHeader } from './authentication';

const base:string[] = ["https://localhost:7239","https://localhost:7239"];

axios.interceptors.response.use(
    (response: AxiosResponse) => response,
    (error: AxiosError) => {
        if (error.message === 'Network Error' && ! error.response) {
            errorMessageStore.Error('network_error_title','network_error_title');
        } else if (error.message !== 'Network Error' && error.response) {
            switch (error.response.status) {
                case 504:
                    errorMessageStore.Error('gateway_timeout_title','gateway_timeout_msg');
                    break;
                case 503:
                    errorMessageStore.Error('service_unavailable_title','service_unavailable_msg');
                    break;
                case 401:
                    errorMessageStore.Error('unauthorized_title','unauthorized_msg');
                    break;
                default:
                    errorMessageStore.Error('internal_server_error_title','internal_server_error_msg');
                    break;
            }
        }

        return Promise.reject(error);
    },
);

const commonAxiosRequests = {
    delNodeRed:(requestUrl:string) => getAuthHeader().then((authResponse) => {
        return axios({
            method: 'delete',
            url: base[1]+requestUrl,
            headers: {
                Authorization: authResponse.Authorization,
            },
        });
    }),
    getNodeRed:(requestUrl: string) => getAuthHeader()
    .then((authResponse) => {
        return axios({
            method: 'get',
            url: base[1]+requestUrl,
            headers: { Authorization: authResponse.Authorization },
        });
    }),
    postNodeRed: (requestUrl: string, requestBody: {}) => getAuthHeader()
        .then((authResponse) => {
            return axios({
                method: 'post',
                url: base[1]+requestUrl,
                headers: {
                    'Authorization': authResponse.Authorization,
                    'Content-Type': 'application/json',
                },
                data: requestBody,
            });
        }),
    get: (requestUrl: string) => getAuthHeader()
        .then((authResponse) => {
            return axios({
                method: 'get',
                url: base[0]+requestUrl,
                headers: { Authorization: authResponse.Authorization },
            });
        }),
    post: (requestUrl: string, requestBody: {}) => getAuthHeader()
        .then((authResponse) => {
            return axios({
                method: 'post',
                url: base[0]+requestUrl,
                headers: {
                    'Authorization': authResponse.Authorization,
                    'Content-Type': 'application/json',
                },
                data: requestBody,
            });
        }),
    put: (requestUrl: string, requestBody: {}) => getAuthHeader()
        .then((authResponse) => {
            return axios({
                method: 'put',
                url: base[0]+requestUrl,
                headers: {
                    'Authorization': authResponse.Authorization,
                    'Content-Type': 'application/json',
                },
                data: requestBody,
            });
        }),
    patch:(requestUrl:string,requestBody:{})=>getAuthHeader()
    .then((authResponse)=>{
        return axios({
            method:'patch',
            url:base[0]+requestUrl,
            headers:{
                'Authorization': authResponse.Authorization,
                'Content-Type': 'application/json'
            },
            data:requestBody
        });
    }),
    del: (requestUrl: string) => getAuthHeader()
        .then((authResponse) => {
            return axios({
                method: 'delete',
                url: base[0]+requestUrl,
                headers: {
                    Authorization: authResponse.Authorization,
                },
            });
        }),
    delWithBody: (requestUrl: string, requestBody: {}) => getAuthHeader()
        .then((authResponse) => {
            return axios({
                method: 'delete',
                url: base[0]+requestUrl,
                headers: {
                    'Authorization': authResponse.Authorization,
                    'Content-Type': 'application/json',
                },
                data: requestBody,
            });
        }),
    getFile: (requestUrl: string) => getAuthHeader()
    .then((authResponse) => {
        return axios({
            method: 'get',
            url: base[0]+requestUrl,
            headers: { Authorization: authResponse.Authorization},
            responseType: 'blob',
        });
    }),
    fetch: (requestUrl: string, requestBody: {}) => getAuthHeader()
        .then((authResponse) => {
            return axios({
                method: 'post',
                url: base[0]+requestUrl,
                headers: {
                    'Authorization': authResponse.Authorization,
                    'Content-Type': 'application/json',
                },
                data: requestBody,
            });
        }),
};

export default commonAxiosRequests;