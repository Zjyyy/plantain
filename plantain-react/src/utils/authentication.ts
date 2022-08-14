// import axios from 'axios';
import { Cookies } from 'react-cookie';
// import commonAxiosRequests from './agent';

export function getAccessToken() {
    const cookies = new Cookies();
    return cookies.get('access_token');
}

export async function getAuthHeader() {
    return { 'Authorization': `Bearer test`};
}