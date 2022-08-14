import { EditOutlined, EllipsisOutlined, SettingOutlined,PlusSquareOutlined,ShareAltOutlined,DeleteOutlined } from '@ant-design/icons';
import { Avatar, Button, Card,Divider, Space, Tooltip } from 'antd';
import React from 'react';
import { useNavigate } from 'react-router-dom';

const { Meta } = Card;

type Props = {
    className?:string,
    style?:React.CSSProperties | undefined
}

const MutimediaMaterials:React.FC<Props> = (props:Props) => {
    const navigate = useNavigate();
    return(
        <div 
        className={`${ props.className ? props.className : ''} multimedia_materials_block`} 
        style={props.style}
        >
            <Divider orientation="right">
                <Space size="large">
                    房间宣传资料
                    <Tooltip placement="top" title="添加资料">
                        <Button type="primary" icon={<PlusSquareOutlined />}/>
                    </Tooltip>
                    
                </Space>
            </Divider>

            <div className='multimedia_card_panel'>
                <Card
                    className='mutimedia_card'
                    cover={
                    <img
                        alt="example"
                        src="https://gw.alipayobjects.com/zos/rmsportal/JiqGstEfoWAOHiTxclqi.png"
                    />
                    }
                    actions={[
                        <Tooltip placement="top" title="创建分享">
                            <ShareAltOutlined key="setting" />
                        </Tooltip>,
                        <Tooltip placement="top" title="编辑">
                            <EditOutlined key="edit" onClick={()=>{navigate('/apartmentAttribute/multimediaMaterialsAttribute')}}/>
                        </Tooltip>,
                        <Tooltip placement="top" title="删除">
                            <DeleteOutlined style={{color:'red'}} key="setting" />
                        </Tooltip>                 
                    ]}
                >
                    <Meta
                    avatar={<Avatar src="https://joeschmoe.io/api/v1/random" />}
                    title="Card title"
                    description="This is the description"
                    />
                </Card>

                <Card
                    className='mutimedia_card'
                    cover={
                    <img
                        alt="example"
                        src="https://gw.alipayobjects.com/zos/rmsportal/JiqGstEfoWAOHiTxclqi.png"
                    />
                    }
                    actions={[
                        <Tooltip placement="top" title="创建分享">
                            <ShareAltOutlined key="setting"/>
                        </Tooltip>,
                        <Tooltip placement="top" title="编辑">
                            <EditOutlined key="edit" />
                        </Tooltip>,
                        <Tooltip placement="top" title="删除">
                            <DeleteOutlined style={{color:'red'}} key="setting" />
                        </Tooltip>                 
                    ]}
                >
                    <Meta
                    avatar={<Avatar src="https://joeschmoe.io/api/v1/random" />}
                    title="Card title"
                    description="This is the description"
                    />
                </Card>
                <Card
                    className='mutimedia_card'
                    cover={
                    <img
                        alt="example"
                        src="https://gw.alipayobjects.com/zos/rmsportal/JiqGstEfoWAOHiTxclqi.png"
                    />
                    }
                    actions={[
                        <Tooltip placement="top" title="创建分享">
                            <ShareAltOutlined key="setting"/>
                        </Tooltip>,
                        <Tooltip placement="top" title="编辑">
                            <EditOutlined key="edit" />
                        </Tooltip>,
                        <Tooltip placement="top" title="删除">
                            <DeleteOutlined style={{color:'red'}} key="setting" />
                        </Tooltip>                 
                    ]}
                >
                    <Meta
                    avatar={<Avatar src="https://joeschmoe.io/api/v1/random" />}
                    title="Card title"
                    description="This is the description"
                    />
                </Card>
            </div>
        </div>
    );
}

export default MutimediaMaterials;