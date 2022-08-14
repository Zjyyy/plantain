import React from 'react'
import ClientResourceStore from './clientResource/ClientResourceStore'
import ClientMessageModalStore from './components/ClientMessageModalStore'
import HouseResourceStore from './housingResource/HouseResourceStore'
import HouseModalStore from './components/HouseModalStore'

export const rootStore = {
    ClientResourceStore:new ClientResourceStore(),
    ClientMessageModalStore:new ClientMessageModalStore(),
    HouseResourceStore:new HouseResourceStore(),
    HouseModalStore:new HouseModalStore()
}