import React from "react";
import { Button,Space,Input,Divider } from "antd";
import MultimediaMaterialsAttributePhoto from './MultimediaMaterialsAttributePhoto'
import MultimediaMaterialsAttributeVideo from "./MultimediaMaterialsAttributeVideo";

const { TextArea } = Input;
const MultimediaMaterialsAttribute:React.FC = () => {
    return(
        <>
        <Button type="primary">保存</Button>
        <div className="multimedia_materials_attribute">
            <div className="multimedia_materials_attribute_left">
                <Divider orientation="left" >基本信息</Divider>
                <Space className="client_message_space">
                    <label>标题:</label>
                    <Input style={{width:150}} placeholder="标题" />
                </Space>
                <Space className="client_message_space">
                    <label>描述:</label>
                    <Input style={{width:150}} placeholder="描述" />
                </Space>
                <Space className="client_message_space">
                    <label>宣传文案:</label>
                    <TextArea style={{width:300}} rows={4} />
                </Space>
            </div>
            <div className="multimedia_materials_attribute_right">
                <Divider orientation="left" >图片</Divider>
                <MultimediaMaterialsAttributePhoto/>

                <Divider orientation="left" >视频</Divider>
                <MultimediaMaterialsAttributeVideo/>
            </div>
        </div>
        </>
    );
}

export default MultimediaMaterialsAttribute;