import React, { useState } from "react";
import ClientMessageModal from "./ClientMessageModal";
import { 
    Select, Space,Input,
    DatePicker,Timeline,
    Divider,Button,Tooltip,Modal  
} from 'antd';
import { 
    ClockCircleOutlined,UserAddOutlined,
    MinusOutlined,ExclamationCircleOutlined
} from '@ant-design/icons';

const { Option } = Select;
const { RangePicker } = DatePicker;
const { confirm } = Modal;

type Props = {
    className?:string,
    style?:React.CSSProperties | undefined;
}

const BasicInformation:React.FC<Props> = (props:Props) => {
    const [clientMessageModelVisiable,setClientMessageModelVisiable] = useState(false);

    const showDeleteConfirm = () => {
        confirm({
          title: '确认移除该租客?',
          icon: <ExclamationCircleOutlined />,
          content: '',
          okText: '确定',
          okType: 'danger',
          cancelText: '取消',
          onOk() {
            console.log('OK');
          },
          onCancel() {
            console.log('Cancel');
          },
        });
      };

    const handleChange = (value: string) => {
        console.log(`selected ${value}`);
    };

    return(
        <div 
        className={`${ props.className ? props.className : '' } basic_information_block`} 
        style={props.style}
        >
            <Divider orientation="left" >房间信息</Divider>
            <Space className="basic_information_space">
                <label>房间所属楼:</label>
                <Select defaultValue="all" style={{ width: 350 }} onChange={handleChange}>n
                    <Option value="all">杭州盈和清机大厦</Option>
                    <Option value="rentedOut">杭州0纪元</Option>
                    <Option value="notRentOut">杭州山水时代</Option>
                </Select>
            </Space>
            <Space className="basic_information_space">
                <label>房间门牌号:</label>
                <Input placeholder="门牌号" />
            </Space>

            <Divider orientation="left" >租客信息</Divider>
                <Tooltip placement="rightTop" title="添加租客">
                    <Button icon={<UserAddOutlined />} size="middle" onClick={()=>{ setClientMessageModelVisiable(true) }} />
                </Tooltip>
                <Space className="basic_information_space">
                    <label>租住人:</label>
                    <Input disabled/>
                    <Tooltip placement="rightTop" title="移除租客">
                        <Button danger shape="circle" size="small" icon={<MinusOutlined />} onClick={showDeleteConfirm}/>
                    </Tooltip>
                </Space>
                <Divider orientation="left" >租约信息</Divider>
                <Space className="basic_information_space">
                    <label>月租金:</label>
                    <Input prefix="￥" suffix="RMB" />
                </Space>
                <Space className="basic_information_space">
                    <label>已交押金:</label>
                    <Input prefix="￥" suffix="RMB" />
                </Space>
                <Space className="basic_information_space">
                    <label>支付方式:</label>
                    <Select defaultValue="1" style={{ width: 350 }} onChange={handleChange}>n
                        <Option value="1">月付</Option>
                        <Option value="2">2月付</Option>
                        <Option value="3">季付</Option>
                        <Option value="5">半年付</Option>
                        <Option value="12">年付</Option>
                    </Select>
                </Space>
                <Space className="basic_information_space">
                    <label>下次收租日期:</label>
                    <DatePicker />
                    
                </Space>
                <Space className="basic_information_space">
                    <label>租约日期:</label>
                    <RangePicker />
                </Space>
                <Divider orientation="left" >租金支付记录</Divider>
                <Space className="basic_information_space">
                    <Timeline mode="left" style={{width:'300px'}}>
                        <Timeline.Item label="2015-09-01" color="green">已交租</Timeline.Item>
                        <Timeline.Item label="2015-09-01 09:12:11" color="green">已交租</Timeline.Item>
                        <Timeline.Item label="2015-09-01 09:12:11" dot={<ClockCircleOutlined className="timeline-clock-icon" />} color="red">交租逾期</Timeline.Item>
                        <Timeline.Item label="2015-09-01 09:12:11">还未交租</Timeline.Item>
                    </Timeline>
                </Space>

            <ClientMessageModal 
                visible={clientMessageModelVisiable}
                onOk={()=>{setClientMessageModelVisiable(false)}}
                onCancel={()=>{setClientMessageModelVisiable(false)}}
            />
        </div>
    );
}

export default BasicInformation;