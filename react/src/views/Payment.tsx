import React from 'react'
import { useLocation } from 'react-router-dom'
import { IContract } from '~sections/CreateContract/CreateContract';
import PaymentSection from '~sections/Payment/PaymentSection'

function PaymentView() {
    const contract = (useLocation().state as IContract);
    return (
        <PaymentSection contract={contract} />
    )
}

export default PaymentView