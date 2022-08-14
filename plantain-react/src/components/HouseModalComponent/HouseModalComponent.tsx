import React, { useEffect } from "react"
import { Modal,Space,Input } from 'antd'
import { useStores,observer } from "src/stores/storeHook"
import style from "src/scss/component/houseModal.module.scss"

const {TextArea} = Input

type Props = {
    houseId?:string,
    visible:boolean,
    onOk?:(...args:any[]) => any,
    onCancel?:(...args:any[]) => any
}

const HouseModalComponent:React.FC<Props> = observer((props:Props) => {
    const store = useStores().HouseModalStore

    useEffect(()=>{
        if(props.houseId != undefined){
            store.loadHouseMessage(props.houseId)
        }
    },[props.visible == true])

    return(
        <Modal
        title="楼信息编辑"
        centered
        visible={props.visible}
        onOk={props.onOk}
        onCancel={props.onCancel}
        width={500}
        >
            <div className={style.content}>
                <Space className={style.container}>
                    <label>楼栋名:</label>
                    <Input 
                    style={{width:250}} 
                    value={store.house.name}
                    onChange={(e)=>store.onChangeName(e.target.value as string)}
                    placeholder="楼栋名" />
                </Space>
                <Space className={style.container}>
                    <label>描述:</label>
                    <TextArea 
                    style={{width:250}} 
                    value={store.house.des}
                    onChange={(e)=>store.onChangeDes(e.target.value as string)}
                    rows={4}/>
                </Space>
            </div>
        </Modal>
    )
})

export default HouseModalComponent