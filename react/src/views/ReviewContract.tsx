import React from 'react'
import { useLocation } from 'react-router-dom'
import { IContract } from '~sections/CreateContract/CreateContract';
import ReviewContractSection from '~sections/ReviewContract';

function ReviewContract() {
    const contract = useLocation().state as IContract;

    return (
        <div>
            <div className="messages">Success message here</div>
            <div className="title">Review escrow Contract #{contract.id}</div>
            <ReviewContractSection contract={contract} />
        </div>
    )
}

export default ReviewContract