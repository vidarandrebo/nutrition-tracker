import {IAssignFromForm} from "./IAssignFromForm.ts";
import {getStringField} from "../components/forms/FormUtils.ts";
import {User} from "./User.ts";
import {v4 as UuidV4} from "uuid";

export class Credentials implements IAssignFromForm {
    email: string;
    password: string;

    constructor() {
        this.email = "";
        this.password = "";
    }

    assignFromFormData(form: FormData) {
        this.email = getStringField(form, "email");
        this.password = getStringField(form, "password");
    }

    loginUser(): User {
        // TODO
        // send login user request to server
        const user = new User();
        user.userId = UuidV4(); //TMP
        user.email = this.email; //TMP
        return user;
    }

    registerUser(): User {
        // TODO
        // send register user request to server
        const user = new User();
        return user;
    }
}