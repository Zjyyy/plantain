import React from "react";
import { Modal,Space,Input } from 'antd';

const {TextArea} = Input;

type Props = {
    visible:boolean,
    onOk?:(...args:any[]) => any,
    onCancel?:(...args:any[]) => any
}

const ClientMessageModel:React.FC<Props> = (props:Props) => {
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
                <Input style={{width:150}} placeholder="姓名" />
            </Space>
            <Space className="basic_information_space">
                <label>身份证:</label>
                <Input style={{width:300}} placeholder="身份证" />
            </Space>
            <Space className="basic_information_space">
                <label>联系方式:</label>
                <Space direction="vertical" size="small">
                    <Input addonBefore="手机号" style={{width:300}} placeholder="手机号" />
                    <Input addonBefore="微信号" style={{width:300,marginTop:'-8px'}} placeholder="微信号" />
                </Space>
            </Space>
            <Space className="basic_information_space">
                <label>备注:</label>
                <TextArea style={{width:300}} rows={4} />
            </Space>
        </div>
      </Modal>
    );
}

export default ClientMessageModel;