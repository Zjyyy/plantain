import React, { useEffect } from "react";
import { Divider,Space,Input,Rate,Timeline, Button, message } from "antd";
import { useParams,useNavigate } from "react-router-dom";
import { useStores,observer } from "src/stores/storeHook";

const {TextArea} = Input;

const ClientMessage:React.FC = observer(() => {
    const { id } = useParams()
    const store = useStores().ClientResourceStore
    const navigate = useNavigate(); 

    useEffect(()=>{
        if(id !== undefined){
            store.loadClientMessage(id)
        }
    },[])

    const onClickSave = async() => {
        if(await store.submitClientMessage()){
            message.success("保存客户成功")
            navigate('/clientList')
        }
    }

    return(
    <>
    <Button type="primary" onClick={onClickSave}>保存</Button>
    <div className="client_message">
        <div className="client_basic_message">
            <Divider orientation="left" >基本信息</Divider>
            <Space className="client_message_space">
                <label>姓名:</label>
                <Input style={{width:150}} placeholder="姓名" 
                onChange={(e) => store.onChangeName(e.target.value as string)} value={store.clientMessage.name}/>
            </Space>
            <Space className="client_message_space">
                <label>身份证:</label>
                <Input style={{width:300}} placeholder="身份证" 
                onChange={(e) => store.onChangeIDNumber(e.target.value as string)} value={store.clientMessage.idNumber}/>
            </Space>
            <Space className="client_message_space">
                <label>联系方式:</label>
                <Space direction="vertical" size="small">
                    <Input addonBefore="手机号" style={{width:300,margin:'-4px 0px'}} 
                    onChange={(e)=>store.onChangePhoneNumber(e.target.value as string)}
                    placeholder="手机号" value={store.clientMessage.phoneNumber}/>
                    <Input addonBefore="微信号" style={{width:300,margin:'-4px 0px'}} 
                    onChange={(e)=>store.onChangeWeChatNumber(e.target.value as string)}
                    placeholder="微信号" value={store.clientMessage.weChatNumber}/>
                </Space>
            </Space>
            <Space className="client_message_space">
                <label>评分:</label>
                <Rate 
                onChange={(val) => store.onChangeRate(val)}
                value={store.clientMessage.rate}/>
            </Space>
            <Space className="client_message_space">
                <label>备注:</label>
                <TextArea style={{width:300}} rows={4} 
                onChange={(e) => store.onChangeDes(e.target.value as string)} value={store.clientMessage.des}/>
            </Space>
        </div>

        <div className="client_history_message">
            <Divider orientation="right" >租住记录</Divider>
            <Timeline mode="left" style={{width:'300px'}}>
                <Timeline.Item label="2015-09-01" color="green">入住</Timeline.Item>
                <Timeline.Item label="2015-10-01" color="green">按时交租</Timeline.Item>
                <Timeline.Item label="2015-11-01" color="red">逾期交租</Timeline.Item>
                <Timeline.Item label="2015-12-01">退租</Timeline.Item>
            </Timeline>
        </div>
    </div>
    </>
    );
})

export default ClientMessage