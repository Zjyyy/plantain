import {Component} from 'react';
import { Menu,Breadcrumb } from 'antd';
import { MailOutlined, AppstoreOutlined, SettingOutlined } from '@ant-design/icons';
import { useRoutes,Link,useLocation } from 'react-router-dom'
import logoUrl from 'src/logo.png';
import { breadcrumbNameMap } from 'src/model/NameMapModel'

const { SubMenu } = Menu;

export const TopNaviBar:React.FC = () =>{
    const location = useLocation();
    const pathSnippets = location.pathname.split('/').filter(i=>i);
    const extraBreadcrumbItems = pathSnippets.map((_, index) => {
        const url = `/${pathSnippets.slice(0, index + 1).join('/')}`;
        return (
        <Breadcrumb.Item key={url}>
            <Link to={url}>{breadcrumbNameMap[url]}</Link>
        </Breadcrumb.Item>
        );
    });
    const breadcrumbItems = [
        <Breadcrumb.Item key="home">
            <Link to="/">Home</Link>
        </Breadcrumb.Item>,
    ].concat(extraBreadcrumbItems);

    return(
        <div className='wfc_top_nav'>
            <img  style={{width:'25px',height:'25px'}} alt="logo" src={logoUrl}></img>
            <Breadcrumb style={{ margin: '5px 0 0 15px' }}>
                {breadcrumbItems}
            </Breadcrumb>
        </div>
    )
}

export default TopNaviBar;