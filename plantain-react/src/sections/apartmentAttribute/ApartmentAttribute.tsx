import { Button } from 'antd';
import * as React from 'react';
import BasicInformation from './BasicInformation';
import MutimediaMaterials from './MultimediaMaterials';

const ApartmentAttribute:React.FC = () => {
    return(
        <>            
        <Button type="primary">保存</Button>
        <div className='apartment_attribute_panel'>
            <BasicInformation className='apartment_basic_information'/>
            <MutimediaMaterials className='apartment_multimedia_materials'/>
        </div>
        </>
    )
}

export default ApartmentAttribute;