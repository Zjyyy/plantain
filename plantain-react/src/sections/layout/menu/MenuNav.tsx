import React,{ useState } from "react";
import { Menu,Button } from 'antd';
import { Link } from 'react-router-dom';
import {
    BankOutlined,
    FormOutlined,
    SettingOutlined,
    DatabaseOutlined,
    InsertRowAboveOutlined,
    OrderedListOutlined,
    LaptopOutlined,
    LeftOutlined,
    RightOutlined,
    DashboardOutlined,
    BookOutlined,
    PieChartOutlined,
    ClusterOutlined,
    ContactsOutlined,
    ShareAltOutlined
  } from '@ant-design/icons';

import type { MenuProps } from 'antd';
type MenuItem = Required<MenuProps>['items'][number];
function getItem(
    label: React.ReactNode,
    key: React.Key,
    icon?: React.ReactNode,
    children?: MenuItem[],
    type?: 'group',
  ): MenuItem {
    return {
      key,
      icon,
      children,
      label,
      type,
    } as MenuItem;
  }
  
  const items: MenuItem[] = [
    getItem(<Link to='/workbench'>面板</Link>, '1', <LaptopOutlined />),
    getItem(<Link to='/housePanel'>房源</Link>, 'sub1', <BankOutlined />),
    getItem(<Link to='/clientList'>客户</Link>, 'clientList', <ContactsOutlined />),
    getItem('财务', 'sub2', <BookOutlined />,[
      getItem(<Link to='/calcTemplate/list'>计算模板列表</Link>, '5', <OrderedListOutlined />),
      getItem(<Link to='/calcTemplate/edit'>计算模板编辑器</Link>, '6', <FormOutlined />),
    ]),
    getItem('报表', 'sub3', <PieChartOutlined />,[
      getItem(<Link to='/monitor/taskState'>计算任务状态监控</Link>, '8', <ClusterOutlined />),
      getItem(<Link to='/monitor/dashboardShow'>Grafana</Link>, '9', <DashboardOutlined />),
    ]),
    getItem('配置', 'sub6', <SettingOutlined />, [
      getItem(<Link to='/system/rtdbconfigure'>RTDB变量配置</Link>, '15',<InsertRowAboveOutlined />),
      getItem(<Link to='/system/settings'>系统配置</Link>, '16',<DatabaseOutlined />)
    ]),
  ];

const MenuNav:React.FC = () => {
    const [collapsed, setCollapsed] = useState(false);

    const toggleCollapsed = () => {
        setCollapsed(!collapsed);
    };
    return(
        <div style={{height:'100%',display:'flex'}}>
            <Menu
                defaultSelectedKeys={['1']}
                defaultOpenKeys={['sub1']}
                mode="inline"
                theme="light"
                inlineCollapsed={collapsed}
                items={items}
            />
            <button onClick={toggleCollapsed} style={{width:'2px',border:0}}>
                {collapsed?<RightOutlined style={{fontSize:'10px'}}/>:<LeftOutlined style={{fontSize:'10px'}}/>}
            </button>
        </div>
    );
}

export default MenuNav;