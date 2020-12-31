import PyPDF2

def check_file_existence(filename: str):
    try:
        open(filename, 'rb')
    except FileNotFoundError:
        print("ERRORE: Il file non esiste")
        exit(2)

def string_to_list(input: str):
    raw_list = input.split(" ")
    normalized_list = []
    for item in raw_list:
        if "-" in item:
            ends = item.split("-")
            for index in range(int(ends[0]), int(ends[1]) + 1):
                # for>if>for? maybe can find something better
                normalized_list.append(index)
        else:
            normalized_list.append(int(item))
    return sorted(normalized_list)
                

def extract_pages(filename: str, pages: list):
    pdf_file = open(filename, 'rb')
    pdf_obj = PyPDF2.PdfFileReader(pdf_file)
    pdfWriter = PyPDF2.PdfFileWriter()
    for page in pages:
        pdfWriter.addPage(pdf_obj.getPage(page - 1))
    
    with open("extracted.pdf", "wb") as f: 
        pdfWriter.write(f)


def extract_pages_option():
    print("# Digitare il nome del file")
    filename = input()
    check_file_existence(filename)
    print("# Digita le singole pagine che vuoi estrarre separate da uno spazio e gli intervalli da un trattino")
    print("Esempio: 1 4 18 25-37 78-90 98")
    page_list = input()
    extract_pages(filename, string_to_list(page_list))
    

def merge_pdfs(pdfs: list):
    # check https://stackoverflow.com/questions/49927338/merge-2-pdf-files-giving-me-an-empty-pdf/49927541#49927541
    merger = PyPDF2.PdfFileMerger()
    for pdf in pdfs:
        merger.append(PyPDF2.PdfFileReader(pdf), 'rb')
    with open("merged.pdf", 'wb') as outfile:
        merger.write(outfile)


def merge_pdfs_option():
    i = 1
    pdfs_list = []
    while True:
        print(f"# Digitare il nome del {i}° file. Premere invio senza inserire nulla per proseguire")
        filename = input()
        if filename == "":
            break
        check_file_existence(filename)
        pdfs_list.append(filename)
        i = i + 1
    merge_pdfs(pdfs_list)


if __name__ == "__main__":
    options = {"1":extract_pages_option,
               "2":merge_pdfs_option, 
               "q":exit}
    while True:
        print("# Seleziona una delle seguenti opzioni:")
        print("1) Estrarre pagine")
        print("2) Unire più PDF")
        print("q) Esci dal programma")
        selection = input()
        if selection not in options:
            print("ERRORE: inserito parametro imprevisto, i possibili parametri erano:", end="")
            for item in options:
                print(f' "{item}"', end="")
            print()
        else:
            options[selection]()
