import { Modal, Space } from 'antd';

export class ErrorMessageStore{
    public Error(title:string,content:string){
        Modal.error({
            title:title,
            content:content
        })
    }
}

const errorMessageStore = new ErrorMessageStore();
export default errorMessageStore;