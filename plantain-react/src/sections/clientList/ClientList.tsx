import { EditOutlined,DeleteOutlined } from '@ant-design/icons';
import { Avatar, Card,Button,Tooltip,message } from 'antd';
import React, { useState,useEffect } from 'react';
import { useNavigate } from "react-router-dom"
import { IClientPanel } from 'src/model/clientResourceModel';
import { useStores,observer } from 'src/stores/storeHook';
import { ClientMessageModalComponent } from 'src/components/ClientMessageModalComponent';

const { Meta } = Card;

const ClitenList:React.FC = observer(()=>{
    const navigate = useNavigate();
    const [visibleAddMessageModal,setVisibleAddMessageModal] = useState(false)
    const store = useStores().ClientResourceStore;
    const storeModal = useStores().ClientMessageModalStore
    useEffect(()=>{
        store.loadClientPanelList()
    },[])

    const onSaveClientMessage = async() => {
        if(await storeModal.submitClientMessage()){
            message.success("添加客户成功")
            store.loadClientPanelList()
        }
        setVisibleAddMessageModal(false)
    }
    const onCancelClientMessage = () => {
        storeModal.initializeClientMessage()
        setVisibleAddMessageModal(false)
    }
    const onClickRemoveClientMessage = async(id:string) =>{
        if(await store.removeClientMessage(id)){
            message.success("删除客户成功")
            store.loadClientPanelList()
        }
    }
    return(
    <>
        <div style={{padding:'5px 5px',width:'100%',background:'#f0f0f0'}}>
            <Button type="primary" onClick={()=>{setVisibleAddMessageModal(true)}}>添加</Button>
        </div>
        <div style={{display:'flex'}}>
        {
            store.clientPanelList.map((item:IClientPanel,key)=>{
                return(
                    <Card
                        key={key}
                        style={{ width: 300, margin:'20px 10px' }}
                        actions={[
                            <Tooltip placement="top" title="编辑">
                                <EditOutlined key="edit" onClick={()=>{navigate('/clientList/clientMessage/'+item.id)}}/>
                            </Tooltip>,

                            <Tooltip placement="top" title="删除">
                                <DeleteOutlined style={{color:'red'}} key="setting" onClick={()=> onClickRemoveClientMessage(item.id)}/>
                            </Tooltip>
                        ]}
                    >
                        <Meta
                            avatar={<Avatar src="https://joeschmoe.io/api/v1/random" />}
                            title={item.name}
                            description={item.des}
                        />
                    </Card>
                )
            })
        }
        </div>
        <ClientMessageModalComponent 
                visible={visibleAddMessageModal}
                onOk={  onSaveClientMessage  }
                onCancel={ onCancelClientMessage }
        />
    </>
    );
})

export default ClitenList;