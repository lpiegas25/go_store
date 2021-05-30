import axios from 'axios';
import React, { useState } from 'react'

export const RoleForm = () => {

    const [formState, setFormState] = useState({
        name: "",
        description: ""
    });

    const { name, description } = formState;

    const handleInputChange = ({ target }: React.ChangeEvent<HTMLInputElement>) => {
        setFormState({
            ...formState,
            [target.name]: target.value
        })
    }

    const endpoint = 'http://localhost:8000/api/v1/roles';

    const handleSubmit = async () => {
        await axios.post(
            endpoint,
            formState,
        );
        resetFormState();
    }

    const resetFormState = () => {
        setFormState({
            name: "",
            description: ""
        });
    }

    return (
        <div>
            <form onSubmit={(e) => { e.preventDefault(); handleSubmit();}}>
                <input
                    type="text"
                    value={name}
                    onChange={handleInputChange}
                    placeholder="name"
                    name='name'
                />
                <input
                    type="text"
                    placeholder="description"
                    onChange={handleInputChange}
                    value={description}
                    name='description'
                />
                <button type='submit'>Save</button>
            </form>
        </div>
    )
}
