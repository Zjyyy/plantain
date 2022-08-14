import React, { useEffect, useLayoutEffect, useState } from "react";
import { EditOutlined,DeleteOutlined } from '@ant-design/icons';
import GridLayout from 'react-grid-layout';
import { useNavigate } from "react-router-dom"
import { Card,Statistic,Space,Button,message,Tooltip } from "antd";
import { Link } from "react-router-dom";
import { 
    BankOutlined,
    AccountBookOutlined 
}from '@ant-design/icons';
import { useStores,observer } from "src/stores/storeHook";
import { HouseModalComponent } from "src/components/HouseModalComponent"

const HouseStatusPanel:React.FC = observer(() => {
    const navigate = useNavigate();
    const store = useStores().HouseResourceStore
    const storeModal = useStores().HouseModalStore
    const [visibleAddHouseModal,setVisibleAddHouseModal] = useState(false)
    
    useEffect(()=>{
        store.loadHouseList()
    },[])

    const onCancelHouse = () => {
        storeModal.initializeHouse()
        setVisibleAddHouseModal(false)
    }
    const onSaveHouse = async () => {
        if(await storeModal.submitHouse()){
            message.success("添加楼栋成功")
            store.loadHouseList()
        }
        setVisibleAddHouseModal(false)
    }

    return(
        <>
        <div style={{padding:'5px 5px',width:'100%',background:'#f0f0f0'}}>
            <Button type="primary" onClick={() => setVisibleAddHouseModal(true)}>添加</Button>
        </div>
        {
            store.layout.length > 0 && 
            (
                <GridLayout layout={store.layout} cols={12} rowHeight={200} width={1800}>
                    {
                        store.housePanelList.map((item,key)=>{
                            return(
                                <Card 
                                key={item.id} 
                                title={item.name} 
                                extra={<Link to={'/house/'+item.id}>More</Link>} 
                                style={{height:400}}>
                                    <Space size="large">
                                        <Statistic title="已租/总数" value={item.rentedNumber} suffix={"/ "+ item.sumNumber} prefix={<BankOutlined />} />
                                        <Statistic title="本月收入" value={item.monthlyIncome} prefix={<AccountBookOutlined />}/>
                                    </Space>
                                </Card>
                            )
                        })
                    }
                </GridLayout>
            )
        }
        <HouseModalComponent 
            visible = {visibleAddHouseModal}
            onOk={ onSaveHouse }
            onCancel={ onCancelHouse }
        />
        </>
    );
})

export default HouseStatusPanel;