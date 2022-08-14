import React, { HTMLAttributes, useEffect, useState } from "react";
import { useRoutes,useLocation } from "react-router-dom";
import { titleNameMap } from 'src/model/NameMapModel'
import routes from 'src/routes/router'

type Props = {
    className?:string,
    style?:React.CSSProperties | undefined;
}
const Container:React.FC<Props> = (props) => {
    const [title,setTitle] = useState('');
    const elements = useRoutes(routes);
    const location = useLocation();
    const pathSnippets = location.pathname.split('/').filter(_=>_);
    useEffect(()=>{
        pathSnippets[pathSnippets.length - 1] != undefined && setTitle(pathSnippets[pathSnippets.length - 1]);
    });
    
    return (
        <div style={props.style}>
            <h1>{titleNameMap[title]}</h1>
            {elements}
        </div>
    );
}

export default Container;