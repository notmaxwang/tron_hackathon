import React, { useState } from 'react';
import './Payment.css'; // Import your CSS file

const PaymentPage = () => {
  const [step, setStep] = useState(1); // Default step is 1

  const handleNextStep = () => {
    setStep(step + 1);
  };

  const handlePrevStep = () => {
    setStep(step - 1);
  };

  return (
    <div className="payment-page">
      {step === 1 && (
        <div className="step-container">
          <h2>Step 1: Review Listing</h2>
          {/* Content for reviewing the listing */}
          <button onClick={handleNextStep}>Next</button>
        </div>
      )}
      {step === 2 && (
        <div className="step-container">
          <h2>Step 2: Payment</h2>
          {/* Content for payment */}
          <button onClick={handlePrevStep}>Previous</button>
          <button onClick={handleNextStep}>Next</button>
        </div>
      )}
      {step === 3 && (
        <div className="step-container">
          <h2>Step 3: Confirmation</h2>
          {/* Content for confirmation */}
          <button onClick={handlePrevStep}>Previous</button>
        </div>
      )}
    </div>
  );
};

export default PaymentPage;

