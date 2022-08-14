export interface IHousePanel {
    id:string,
    name:string,
    des:string,
    rentedNumber:number,
    sumNumber:number,
    monthlyIncome:number
}

export interface IHouse extends IHousePanel{}

export interface IApartmentPanel {
    id:string,
    roomBelongsHouseId:string,
    apartmentName:string
    monthlyRent:number,
    modeOfPayment:string,
    guaranteeDeposit:number
    collectRentDate:string,
    leaseTime:string
}