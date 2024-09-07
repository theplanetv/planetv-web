CREATE TRIGGER prevent_before_insert_blogtagfile
BEFORE INSERT ON public.blogtagfile
FOR EACH ROW
EXECUTE FUNCTION prevent_if_same_blogtagfile();

CREATE TRIGGER prevent_before_update_blogtagfile
BEFORE UPDATE ON public.blogtagfile
FOR EACH ROW
EXECUTE FUNCTION prevent_if_same_blogtagfile();
