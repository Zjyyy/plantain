import React from "react";
import GridLayout from 'react-grid-layout';
import { Card,Statistic,Space,Button,Select,message } from "antd";
import { Link,useNavigate,useParams } from "react-router-dom";
import { 
    BankOutlined,
    AccountBookOutlined 
}from '@ant-design/icons';
import { HouseModalComponent } from "src/components/HouseModalComponent"
import { useState,useEffect } from 'react';
import { useStores,observer } from "src/stores/storeHook";

const { Option } = Select;
const { Countdown } = Statistic;
const deadline = Date.now() + 1000 * 60 * 60 * 24 * 2 + 1000 * 30; // Moment is also OK

const layout = [
    {i:'dataExtraction',x:0,y:0,w:2,h:1,isResizable:false},
    {i:'calcTask',x:2,y:0,w:2,h:1,isResizable:false},
    {i:'taskState',x:4,y:0,w:2,h:1,isResizable:false}
]
const ApartmentStatusPanel:React.FC = observer(() => {
    const navigate = useNavigate();
    const {id} = useParams()
    const [visibleAddHouseModal,setVisibleAddHouseModal] = useState(false)
    const storeModal = useStores().HouseModalStore
    const store = useStores().HouseResourceStore
    const onCancelHouse = () => {
        storeModal.initializeHouse()
        setVisibleAddHouseModal(false)
    }
    const onSaveHouse = async () => {
        if(await storeModal.patchHouseMessage(id as string)){
            message.success("修改楼栋成功")
        }
        setVisibleAddHouseModal(false)
    }
    const onDeleteHoue = async() => {
        if(await store.removeHouse(id as string)){
            message.success("删除楼栋成功")
            navigate("/housePanel")
        }
    }
    const handleChange = (value: string) => {
        console.log(`selected ${value}`);
    };

    useEffect(()=>{
        store.loadApartmentList()
    },[])
    return(
    <>
        <Space size={"large"} style={{padding:'5px 5px',width:'100%',background:'#f0f0f0'}}>
            <Button onClick={() => setVisibleAddHouseModal(true)}>编辑楼信息</Button>
            <Button danger onClick={onDeleteHoue}>删除楼</Button>
            <Button type="primary">添加公寓</Button>
            <Space size={"small"}>
                <label>筛选:</label>
                <Select defaultValue="all" style={{ width: 120 }} onChange={handleChange}>n
                    <Option value="all">全部</Option>
                    <Option value="rentedOut">已出租</Option>
                    <Option value="notRentOut">未出租</Option>
                    <Option value="rentIsOverdue">租金逾期</Option>
                    <Option value="expiringLease">即将到期</Option>
                </Select>
            </Space>
        </Space>
        <GridLayout layout={store.apartmentLayout} cols={12} rowHeight={230} width={1800}>
            {/* <Card key="dataExtraction" title="101" extra={<Link to='/apartmentAttribute'>More</Link>}>
                <Space direction="vertical" size={"large"}>
                    <Countdown valueStyle={{fontSize:'17px'}} title="剩余租期" value={deadline} format="D 天 H 时 m 分 s 秒" />
                    <Space size="large">
                        <Statistic title="月租" value={200} valueStyle={{fontSize:'17px'}}/>
                        <Statistic title="状态" value={"已出租"} valueStyle={{fontSize:'17px'}}/>
                        <Statistic title="收租日期" value={"2022/6/26"} valueStyle={{fontSize:'17px'}}/>
                    </Space>
                </Space>
            </Card> */}
            {
                store.apartmentPanelList.map((item,key)=>{
                    return(
                        <Card key={item.id} title={item.apartmentName} extra={<Link to='/apartmentAttribute'>More</Link>}>
                            <Space direction="vertical" size={"large"}>
                                <Countdown valueStyle={{fontSize:'17px'}} title="剩余租期" value={item.leaseTime} format="D 天 H 时 m 分 s 秒" />
                                <Space size="large">
                                    <Statistic title="月租" value={item.monthlyRent} valueStyle={{fontSize:'17px'}}/>
                                    <Statistic title="状态" value={""} valueStyle={{fontSize:'17px'}}/>
                                    <Statistic title="收租日期" value={item.collectRentDate} valueStyle={{fontSize:'17px'}}/>
                                </Space>
                            </Space>
                        </Card>
                    )
                })
            }
        </GridLayout>

        <HouseModalComponent 
            houseId = {id}
            visible = {visibleAddHouseModal}
            onOk={ onSaveHouse }
            onCancel={ onCancelHouse }
        />
    </>
    );
})

export default ApartmentStatusPanel;