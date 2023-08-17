import React, { useState } from "react";
import { Link } from "react-router-dom";
import TokenSelector, { IToken } from "~baseComponents/TokenSelector";
import Tips from "~sections/Tips";
import Typography from "~baseComponents/Typography";
import AddConditions, { ICondition } from "./AddConditions";

// Assets
import randomCubes from "~assets/random-cubes.webp";
import Button from "~baseComponents/Button";
import SecondaryButton from "~baseComponents/SecondaryButton";
import BaseModal from "~baseComponents/BaseModal/Index";
import TokenItem from "~baseComponents/TokenItem";
import Card from "~baseComponents/Card";
import ContentContainer from "~layouts/ContentContainer";

export interface IContract {
  initiatorCoins: IToken;
  fulfillerCoins: IToken;
  conditions?: ICondition[];
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
  const [selectedWantedAmount, setSelectedWantedAmount] = useState<number>(0);
  const [selectedTipAmount, setSelectedTipAmount] = useState<number>(0);
  const [modalToOpen, setModalToOpen] = useState<Modals | undefined>();
  const [selectedOwnToken, setSelectedOwnToken] = useState<IToken | undefined>(
    contract?.initiatorCoins,
  );
  console.log("selectedOwnToken", selectedOwnToken);
  const [conditions, setConditions] = useState<ICondition[]>(
    contract?.conditions ?? [],
  );
  const [selectedWantedToken, setSelectedWantedToken] = useState<
    IToken | undefined
  >(contract?.fulfillerCoins);
  const [selectedTokenTips, setSelectedTokenTips] = useState<
    IToken | undefined
  >(contract?.tips);

  console.log("conditions", conditions);

  const handleSaving = (t: IToken | undefined) => {
    switch (modalToOpen) {
      case Modals.Own:
        setSelectedOwnToken(t);
        break;
      case Modals.Wanted:
        setSelectedWantedToken(t);
        break;
      case Modals.Tips:
        setSelectedTokenTips(t);
        break;
      default:
        break;
    }
    setModalToOpen(undefined);
  };

  const handleSelectedAmountChange = (amount: number) => {
    const newSelectedOwnToken: IToken = {
      name: selectedOwnToken?.name || "",
      display: selectedOwnToken?.display || "",
      amount: selectedOwnToken?.amount || 0,
      selectedAmount: amount,
      denom: selectedOwnToken?.denom || "",
      chain_name: selectedOwnToken?.chain_name || "",
      logos: selectedOwnToken?.logos || {
        svg: "",
        png: "",
      },
    };

    setSelectedOwnToken(newSelectedOwnToken);
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
        showOwnedToken = true;
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
        handleClose={() => setModalToOpen(undefined)}
      />
    );
  };

  return (
    <div>
      <img
        src={randomCubes}
        alt="Dredd-Secure"
        className="object-cover absolute z-0 top-32 right-32 drop-shadow-lightOrange opacity-80"
        loading="lazy"
      />
      <div className="relative min-h-screen w-full pt-32">
        <Link to="/dashboard">
          <Typography
            variant="body"
            className="font-revalia text-orange pl-4 xl:pl-16"
          >
            {"< GO BACK"}
          </Typography>
        </Link>
        <ContentContainer className="pt-32 max-w-6xl">
          <div>
            <span className="overheader">
              <Typography variant="small" className="text-white-500">
                STEP 1
              </Typography>
            </span>
            <div className="title-2">
              <Typography variant="h5" className="font-revalia pb-4">
                Create Contract
              </Typography>
            </div>
            {/* <div className="bg-gray rounded-3xl border-[1px] border-white-200"> */}
            <Card progress="25">
              <div className="p-4 md:p-8">
                <AddConditions
                  conditions={conditions}
                  setConditions={setConditions}
                />
                <div className="assets-management">
                  <div className="py-4">
                    <Typography variant="h6" className="font-revalia">
                      Choose Assets for Exchange
                      <span className="text-orange">*</span>
                    </Typography>
                    <Typography variant="body-small" className="text-white-500">
                      To complete this escrow, you must choose an asset you want
                      to give and an asset to receive
                    </Typography>
                  </div>
                  <div className="flex w-full gap-4">
                    <div className="w-6/12 flex flex-col gap-2">
                      <div className="sub-subtitle">
                        <Typography variant="body-small">
                          Select Your Assets
                        </Typography>
                      </div>
                      {selectedOwnToken ? (
                        <TokenItem
                          token={selectedOwnToken}
                          showAmount={false}
                          selected={true}
                          input={true}
                          selectedAmount={selectedOwnToken.selectedAmount}
                          setSelectedAmount={handleSelectedAmountChange}
                          className=""
                          onClick={() => setModalToOpen(Modals.Own)}
                        />
                      ) : (
                        <SecondaryButton
                          text="Select Token"
                          orangeText
                          onClick={() => setModalToOpen(Modals.Own)}
                        />
                      )}
                    </div>
                    <div className="w-6/12 flex flex-col gap-2">
                      <div className="sub-subtitle">
                        <Typography variant="body-small">
                          Asset you want to receive
                        </Typography>
                      </div>
                      {selectedWantedToken ? (
                        <TokenItem
                          token={selectedWantedToken}
                          showAmount={false}
                          selected={true}
                          input={true}
                          selectedAmount={selectedWantedAmount}
                          setSelectedAmount={setSelectedWantedAmount}
                          noMax
                          className=""
                          onClick={() => setModalToOpen(Modals.Wanted)}
                        />
                      ) : (
                        <SecondaryButton
                          text="Select Token"
                          orangeText
                          onClick={() => setModalToOpen(Modals.Wanted)}
                        />
                      )}
                    </div>
                  </div>
                </div>
              </div>
              <Tips
                token={selectedTokenTips}
                onClick={() => setModalToOpen(Modals.Tips)}
                selectedAmount={selectedTipAmount}
                setSelectedAmount={setSelectedTipAmount}
              />
            </Card>
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
        </ContentContainer>
      </div>
      <BaseModal
        open={modalToOpen !== undefined}
        handleClose={() => setModalToOpen(undefined)}
      >
        {displayModal()}
      </BaseModal>
    </div>
  );
};

export default CreateContract;
