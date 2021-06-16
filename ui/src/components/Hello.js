import React from 'react'
import { useFetch } from '../hooks/useFetch'

export const Hello = () => {

       const {loading, data} = useFetch(`http://localhost:8000/api/v1/roles`);
       
       const { roles } = !!data && data;

    return (
        <>
          <h1>Helloo</h1>

          { loading 
            ? <span className='alert alert-primary'>Cargando...</span>  
            : <h2>{roles.map(({id, name, description}) => {
                return (
                <div key={id} className='mb-5'>
                    <h1>Id: {id}</h1>
                    <h1>Name: {name}</h1>
                    <h2>Description: {description}</h2>
                </div>
                );
            })}</h2>}
           
        </>
    )
}
