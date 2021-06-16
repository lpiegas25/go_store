import axios from 'axios';
import { useState, useEffect, useRef } from 'react';


export const useFetch = ( url ) => {
    
    const isMounted = useRef(true);
    const [state, setState] = useState({ data: null, loading: true, error: null });

    useEffect( () => {
        return () => {
            isMounted.current = false;
        }
    }, [])


    useEffect( () => {

        setState({ data: null, loading: true, error: null });

        const getData = async(url) => {
            const {data} = await axios.get(url);
            setState({
                loading: false,
                error: null,
                data: data,
            });
        }
        getData(url);

    },[url])

    return state;
}
