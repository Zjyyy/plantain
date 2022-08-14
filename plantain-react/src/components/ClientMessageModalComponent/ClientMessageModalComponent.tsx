import React from "react";
import { Modal,Space,Input } from 'antd';
import { useStores,observer } from "src/stores/storeHook";

const {TextArea} = Input;

type Props = {
    visible:boolean,
    onOk?:(...args:any[]) => any,
    onCancel?:(...args:any[]) => any
}

const ClientMessageModalComponent:React.FC<Props> = observer((props:Props) => {
    const store = useStores().ClientMessageModalStore
    return(
        <Modal
        title="租客信息编辑"
        centered
        visible={props.visible}
        onOk={props.onOk}
        onCancel={props.onCancel}
        width={500}
      >
        <div className="basic_information_block">
            <Space className="basic_information_space">
                <label>姓名:</label>
                <Input 
                style={{width:150}} 
                placeholder="姓名" 
                onChange={(e)=> store.onChangeName(e.target.value as string)}
                value={store.clientMessage.name}/>
            </Space>
            <Space className="basic_information_space">
                <label>身份证:</label>
                <Input 
                style={{width:300}} 
                placeholder="身份证" 
                onChange={(e)=> store.onChangeIDNumber(e.target.value as string)}
                value={store.clientMessage.idNumber}/>
            </Space>
            <Space className="basic_information_space">
                <label>联系方式:</label>
                <Space direction="vertical" size="small">
                    <Input addonBefore="手机号" style={{width:300}} 
                    placeholder="手机号" onChange={(e)=>store.onChangePhoneNumber(e.target.value as string)}
                    value={store.clientMessage.phoneNumber}/>
                    <Input addonBefore="微信号" style={{width:300,marginTop:'-8px'}} 
                    placeholder="微信号" onChange={(e)=>store.onChangeWeChatNumber(e.target.value as string)}
                    value={store.clientMessage.weChatNumber}/>
                </Space>
            </Space>
            <Space className="basic_information_space">
                <label>备注:</label>
                <TextArea style={{width:300}} 
                onChange={(e)=>store.onChangeDes(e.target.value as string)}
                rows={4} value={store.clientMessage.des}/>
            </Space>
        </div>
      </Modal>
    );
})

export default ClientMessageModalComponent;