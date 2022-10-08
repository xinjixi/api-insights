import { useState } from 'react';
import Button from '@mui/material/Button';
import Dialog from '@mui/material/Dialog';
import FieldItem from '../../../../components/FieldItem/FieldItem';
import DialogBody from '../../../../components/Dialog/DialogBody/DialogBody';
import DialogActions from '../../../../components/Dialog/DialogActions/DialogActions';
import DialogTitle from '../../../../components/Dialog/DialogTitle/DialogTitle';
import { OrganizationData } from '../../../../query/organization';
import OrganizationDropDown from '../OrganizationDropDown/OrganizationDropDown';
import ProductTagDropDown from '../ProductTagDropDown/ProductTagDropDown';
import BusyIcon from '../../../../components/BusyIcon/BusyIcon';
import IconButton from '../../../../components/IconButton/IconButton';
import './EditServiceDialog.scss';
import { ServiceData } from '../../../../query/service';

export type EditingEvent = {
  title: string;
  productTag: string;
  organizationId: string;
  nameId: string;
  description: string;
  visibility: string;
  contactName: string;
  email: string;
  url: string;
};

type Props = {
  service?: ServiceData.Service;
  open: boolean;
  busy: boolean;
  authEnabled: boolean;
  organizationList: OrganizationData.Organization[];
  productTagList: string[];
  handleClose: () => void;
  onEditing: (data: EditingEvent) => void;
};

/**
 * For inputing details for a new service, or editing info of an existing service.
 */
export default function EditServiceDialog(props: Props) {
  const service: ServiceData.Service = (props.service || {
    id: '',
    name_id: '',
    created_at: '',
    updated_at: '',
    title: '',
    product_tag: '',
    description: '',
    organization_id: '',
    visibility: 'public',
    contact: {
      name: '',
      url: '',
      email: '',
    },
  });

  const { contact } = service;

  const [title, setTitle] = useState(service.title);
  const [productTag, setProductTag] = useState(service.product_tag);
  const [organization, setOrganization] = useState(service.organization_id);
  const [description, setDescription] = useState(service.description);
  const [visibilityPrivate, setVisibilityPrivate] = useState(service.visibility === 'private');
  const visibility = (visibilityPrivate) ? 'private' : 'public';
  const [contactName, setContactName] = useState(contact?.name || '');
  const [email, setEmail] = useState(contact?.email || '');
  const [url, setUrl] = useState(contact.url || '');
  const trimmedTitle = title.trim();
  const trimmedProductTag = productTag.trim();
  const trimmedOrganization = organization.trim();
  const trimmedNameId = title.trim().replaceAll(' ', '_').toLowerCase();

  const onEditing = () => {
    if (props.onEditing) {
      props.onEditing({
        title: trimmedTitle,
        productTag: trimmedProductTag,
        organizationId: trimmedOrganization,
        nameId: trimmedNameId,
        description: description.trim(),
        visibility,
        contactName: contactName.trim(),
        email: email.trim(),
        url: url.trim(),
      });
    }
  };

  const isCreateNew = !props.service;
  const dialogTitle = isCreateNew ? 'Create New Service' : 'Edit Service Details';
  const blockTitle = isCreateNew ? "Let's set up your service. This will be visible on the service tile."
    : "Let's update the service. This will be visible on the service tile.";
  const editAction = isCreateNew ? 'Create' : 'Update';
  const invalidInputs = !trimmedTitle || !trimmedProductTag || !trimmedOrganization;
  const onClose = props.busy ? null : props.handleClose;

  const visibilityField = props.authEnabled && (
    <FieldItem label="Private Visibility">
      <input
        checked={visibilityPrivate}
        className="checkbox"
        type="checkbox"
        onChange={(e) => {
          setVisibilityPrivate(e.target.checked);
        }}
      />
    </FieldItem>
  );

  return (
    <Dialog
      className="edit-service-dialog"
      open={props.open}
      fullWidth
      maxWidth="md"
    >
      <DialogTitle onClose={onClose}>{dialogTitle}</DialogTitle>
      <DialogBody className="edit-service-dialog-body">
        <div className="light-bg">
          <div className="block-title">{blockTitle}</div>
          <FieldItem label="What is this service called?">
            <input
              value={title}
              onChange={(e) => {
                setTitle(e.target.value);
              }}
            />
          </FieldItem>
          <FieldItem label="Enter Service description">
            <input
              value={description}
              onChange={(e) => {
                setDescription(e.target.value);
              }}
            />
          </FieldItem>
          <FieldItem label="Organization">
            <OrganizationDropDown
              editBoxReadOnly={props.authEnabled}
              value={organization}
              list={props.organizationList}
              onChange={(e) => {
                setOrganization(e);
              }}
            />
          </FieldItem>
          <FieldItem label="Product Tag">
            <ProductTagDropDown
              value={productTag}
              list={props.productTagList}
              onChange={(e) => {
                setProductTag(e);
              }}
            />
          </FieldItem>
          {visibilityField}
        </div>
        <div className="white-bg">
          <div className="block-title">
            Who is the assigned contact for this service?
          </div>
          <FieldItem label="Contact Name">
            <input
              value={contactName}
              onChange={(e) => {
                setContactName(e.target.value);
              }}
            />
          </FieldItem>
          <FieldItem label="Email">
            <input
              value={email}
              onChange={(e) => {
                setEmail(e.target.value);
              }}
            />
          </FieldItem>
          <FieldItem label="URL">
            <input
              value={url}
              onChange={(e) => {
                setUrl(e.target.value);
              }}
            />
          </FieldItem>
        </div>
      </DialogBody>
      <DialogActions>
        <Button onClick={onClose}>Cancel</Button>
        <IconButton
          icon={props.busy && <BusyIcon busy />}
          disabled={invalidInputs || props.busy}
          onClick={onEditing}
        >
          {editAction}
        </IconButton>
      </DialogActions>
    </Dialog>
  );
}
