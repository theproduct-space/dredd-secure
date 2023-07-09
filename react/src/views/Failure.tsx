import React from 'react'
import FailureSection, { FailureSectionProps } from '~sections/Status/Failure'

function Failure(props: FailureSectionProps) {
    return (
        <FailureSection {...props} />
    )
}

export default Failure