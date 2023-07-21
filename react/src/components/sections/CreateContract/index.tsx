import React, { useState } from "react";
import { Link } from "react-router-dom";
import TokenElement from "~baseComponents/TokenElement";
import TokenSelector, { IToken } from "~baseComponents/TokenSelector";
import Account from "~sections/Account";
import { ConditionTypes } from "~sections/ReviewContractSection";
import Tips from "~sections/Tips";
import Typography from "~baseComponents/Typography";
import ContentContainer from "~layouts/ContentContainer";
import ModalContainer from "~layouts/ModalContainer";
import Condition from "~sections/Condition";

// Assets
import randomCubes from "~assets/random-cubes.webp";
import Button from "~baseComponents/Button";
import SecondaryButton from "~baseComponents/SecondaryButton";
import { Modal } from "@mui/base";
import BaseModal from "~baseComponents/BaseModal/Index";

export interface ICondition {
  type: string;
  prop: string;
  input?: string;
}

export interface IContract {
  initiatorCoins: IToken;
  fulfillerCoins: IToken;
  conditions?: { condition: ICondition; value: string }[];
  tips?: IToken;
  status?: string;
  id?: string;
}

interface CreateContractProps {
  contract: IContract | undefined;
}

const CreateContract = (props: CreateContractProps) => {
  const { contract } = props;

  enum Modals {
    Own,
    Wanted,
    Tips,
  }
  const [modalToOpen, setModalToOpen] = useState<Modals | undefined>();
  const [selectedOwnToken, setSelectedOwnToken] = useState<IToken | undefined>(
    contract?.initiatorCoins,
  );
  const [selectedWantedToken, setSelectedWantedToken] = useState<
    IToken | undefined
  >(contract?.fulfillerCoins);
  const [selectedTokenTips, setSelectedTokenTips] = useState<
    IToken | undefined
  >(contract?.tips);

  const handleSaving = (t: IToken | undefined) => {
    let func;
    switch (modalToOpen) {
      case Modals.Own:
        func = setSelectedOwnToken;
        break;
      case Modals.Wanted:
        func = setSelectedWantedToken;
        break;
      default:
        func = setSelectedTokenTips;
        break;
    }

    func(t);
    setModalToOpen(undefined);
  };

  const displayModal = () => {
    let modal;
    let showOwnedToken = false;
    switch (modalToOpen) {
      case Modals.Own:
        modal = selectedOwnToken;
        showOwnedToken = true;
        break;
      case Modals.Wanted:
        modal = selectedWantedToken;
        break;
      case Modals.Tips:
        modal = selectedTokenTips;
        break;
      default:
        modal = null;
        break;
    }

    return (
      <TokenSelector
        selectedToken={modal}
        onSave={handleSaving}
        ownedToken={showOwnedToken}
      />
    );
  };

  const displayConditionTypes = () => {
    return ConditionTypes.map((condition) => {
      return (
        <option key={condition.type} value={condition.type}>
          {condition.type}
        </option>
      );
    });
  };

  const [conditions, setConditions] = useState<
    { condition: ICondition; value: string }[]
  >(contract?.conditions ?? []);

  const handleAddNewEmptyCondition = () => {
    const array = [...conditions].concat({
      condition: ConditionTypes[0],
      value: "",
    });
    setConditions(array);
  };

  const handleRemoveCondition = (id: number) => {
    const array = conditions.slice(0, id).concat(conditions.slice(id + 1));
    setConditions(array);
  };

  const handleChangeConditionValue = (
    e: React.ChangeEvent<HTMLInputElement>,
    id: number,
  ) => {
    const array = [...conditions];
    array[id].value = e.target.value;
    setConditions(array);
  };

  const handleChangeCondition = (
    e: React.ChangeEvent<HTMLSelectElement>,
    id: number,
  ) => {
    const array = [...conditions];
    array[id].condition =
      ConditionTypes.find((element) => element.type === e.target.value) ??
      ConditionTypes[0];
    setConditions(array);
  };

  return (
    <ContentContainer>
      <img
        src={randomCubes}
        alt="Dredd-Secure"
        className="object-cover absolute z-0 top-32 drop-shadow-lightOrange opacity-80"
        loading="lazy"
      />
      <div className="relative min-h-screen w-full pt-32">
        <Link to="/dashboard">
          <Typography variant="body" className="font-revalia text-orange">
            {"< GO BACK"}
          </Typography>
        </Link>
        <ModalContainer className="pt-32">
          <div>
            <span className="overheader">
              <Typography variant="small" className="text-white-500">
                STEP 1
              </Typography>
            </span>
            <div className="title-2">
              <Typography variant="h5" className="font-revalia pb-4">
                Create Contract <Account />
              </Typography>
            </div>
            <div className="bg-gray rounded-3xl border-[1px] border-white-200">
              <div className="p-8">
                <div className="conditions-management">
                  <div className="subtitle">
                    <Typography variant="body-small" className="font-revalia">
                      Add Conditions<span className="text-orange">*</span>
                    </Typography>
                  </div>
                  {conditions.map((condition, index) => (
                    <div className="condition" key={`add-condition-${index}`}>
                      <Condition
                        condition={condition}
                        index={index}
                        handleChangeCondition={handleChangeCondition}
                        handleRemoveCondition={handleRemoveCondition}
                        displayConditionTypes={displayConditionTypes}
                        className="pb-2"
                      >
                        Condition #{index + 1}
                      </Condition>
                    </div>
                  ))}
                  <button
                    className="add-condition"
                    onClick={handleAddNewEmptyCondition}
                  >
                    <Typography
                      variant="body-small"
                      className="font-revalia text-orange py-4"
                    >
                      + Add Condition
                    </Typography>
                  </button>
                </div>
                <div className="assets-management">
                  <div className="py-4">
                    <Typography variant="body-small" className="font-revalia">
                      Choose Assets for Exchange
                      <span className="text-orange">*</span>
                    </Typography>
                    <Typography variant="body-small" className="text-white-500">
                      To complete this escrow, you must choose an asset you want
                      to give and an asset to receive
                    </Typography>
                  </div>
                  <div className="flex w-full">
                    <div className="w-6/12 flex flex-col gap-2">
                      <div className="sub-subtitle">
                        <Typography variant="body-small">
                          Select Your Assets:
                        </Typography>
                      </div>
                      <SecondaryButton
                        text="Select Token"
                        orangeText
                        onClick={() => setModalToOpen(Modals.Own)}
                      />
                      <SecondaryButton
                        text="Select NFT"
                        orangeText
                        onClick={() => setModalToOpen(Modals.Own)}
                      />
                    </div>
                    <div className="w-6/12 flex flex-col gap-2">
                      <div className="sub-subtitle">
                        <Typography variant="body-small">
                          Asset you want to receive:
                        </Typography>
                      </div>
                      <SecondaryButton
                        text="Select Token"
                        orangeText
                        onClick={() => setModalToOpen(Modals.Wanted)}
                      />
                      <SecondaryButton
                        text="Select NFT"
                        orangeText
                        onClick={() => setModalToOpen(Modals.Wanted)}
                      />
                    </div>
                  </div>
                </div>
              </div>
              <Tips
                selectedToken={selectedTokenTips}
                onClick={() => setModalToOpen(Modals.Tips)}
              />
            </div>
          </div>
          <div className="flex justify-end py-8">
            <Link
              to={"/escrow/pay"}
              state={{
                initiatorCoins: selectedOwnToken,
                fulfillerCoins: selectedWantedToken,
                conditions: conditions,
              }}
            >
              <Button text="Review Contract" className="capitalize" />
            </Link>
          </div>
        </ModalContainer>
      </div>
      <BaseModal
        open={modalToOpen !== undefined}
        handleClose={() => setModalToOpen(undefined)}
        title="Select Token"
      >
        {displayModal()}
      </BaseModal>
    </ContentContainer>
  );
};

export default CreateContract;
