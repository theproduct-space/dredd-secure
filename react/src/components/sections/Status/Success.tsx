import React from 'react'
import { Link } from 'react-router-dom';

export interface SuccessSectionProps {
    successTitle: string;
    successBody?: string;
    continueButton?: JSX.Element;
}

function SuccessSection(props: SuccessSectionProps) {
    const { successTitle, successBody, continueButton } = props;

    return (
        <div className="card">
            <div className="card-header">Complete</div>
            <div className="card-title centered">{successTitle}</div>
            <div className="card-body">{successBody}</div>
            <div>
                <Link to="/app">Return Home</Link>
                {continueButton}
            </div>
        </div>
    )
}

export default SuccessSection